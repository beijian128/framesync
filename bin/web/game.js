// 游戏配置
const config = {
    canvasWidth: 1600,
    canvasHeight: 900,
    playerSize: 20,
    bulletSize: 8,
    bulletSpeed: 10,
    playerSpeed: 5
};

// 游戏状态
const gameState = {
    players: {},
    bullets: {},
    myPlayerId: null,
    lastUpdateTime: 0,
    inputState: {
        up: false,
        down: false,
        left: false,
        right: false,
        mouseX: 0,
        mouseY: 0,
        shooting: false
    },
    pendingInputs: []
};

// 初始化Canvas
const canvas = document.getElementById('gameCanvas');
canvas.width = config.canvasWidth;
canvas.height = config.canvasHeight;
const ctx = canvas.getContext('2d');

// WebSocket连接
const ws = new WebSocket('ws://localhost:8002');

ws.binaryType = 'arraybuffer';

ws.onopen = () => {
    console.log('Connected to server');
};

ws.onmessage = (event) => {
    const data = new DataView(event.data);

    // 读取2字节的数据包大小
    const packetSize = data.getUint16(0, true);

    // 读取4字节的协议号
    const protocolId = data.getUint32(2, true);

    // 读取JSON数据
    const jsonStr = new TextDecoder().decode(data.buffer.slice(6, 6 + packetSize));
    const message = JSON.parse(jsonStr);

    switch (protocolId) {
        case 1: // 初始化消息
            handleInitMessage(message);
            break;
        case 2: // 游戏状态更新
            handleGameStateUpdate(message);
            break;
    }
};

ws.onclose = () => {
    console.log('Disconnected from server');
};

ws.onerror = (error) => {
    console.error('WebSocket error:', error);
};

// 处理初始化消息
function handleInitMessage(message) {
    gameState.myPlayerId = message.playerId;
    console.log(`My player ID: ${gameState.myPlayerId}`);
}

// 处理游戏状态更新
function handleGameStateUpdate(message) {
    // 更新玩家
    gameState.players = message.players;

    // 更新子弹
    gameState.bullets = message.bullets;

    // 更新服务器确认的输入序列号
    if (message.processedInputs && message.processedInputs.length > 0) {
        gameState.pendingInputs = gameState.pendingInputs.filter(
            input => input.sequence > message.processedInputs[message.processedInputs.length - 1]
        );
    }

    // 渲染游戏
    render();
}

// 发送输入状态
let sequenceNumber = 0;
function sendInputState() {
    if (!gameState.myPlayerId) return;

    const input = {
        sequence: ++sequenceNumber,
        up: gameState.inputState.up,
        down: gameState.inputState.down,
        left: gameState.inputState.left,
        right: gameState.inputState.right,
        mouseX: gameState.inputState.mouseX,
        mouseY: gameState.inputState.mouseY,
        shooting: gameState.inputState.shooting,
        timestamp: Date.now()
    };

    gameState.pendingInputs.push(input);

    // 准备数据包
    const jsonStr = JSON.stringify(input);
    const encoder = new TextEncoder();
    const jsonData = encoder.encode(jsonStr);

    // 创建ArrayBuffer: 2字节大小 + 4字节协议号 + JSON数据
    const buffer = new ArrayBuffer(2 + 4 + jsonData.length);
    const dataView = new DataView(buffer);

    // 写入2字节的数据包大小
    dataView.setUint16(0, jsonData.length, true);

    // 写入4字节的协议号 (3表示客户端输入)
    dataView.setUint32(2, 3, true);

    // 写入JSON数据
    const uint8Array = new Uint8Array(buffer);
    uint8Array.set(jsonData, 6);

    // 发送数据
    ws.send(buffer);

    // 重置射击状态
    gameState.inputState.shooting = false;
}

// 渲染游戏
function render() {
    // 清空画布
    ctx.fillStyle = '#222';
    ctx.fillRect(0, 0, config.canvasWidth, config.canvasHeight);

    // 绘制玩家
    for (const playerId in gameState.players) {
        const player = gameState.players[playerId];

        // 绘制玩家
        ctx.fillStyle = playerId === gameState.myPlayerId ? '#00ff00' : '#ff0000';
        ctx.beginPath();
        ctx.arc(player.x, player.y, config.playerSize, 0, Math.PI * 2);
        ctx.fill();

        // 绘制玩家方向指示器
        ctx.strokeStyle = '#ffffff';
        ctx.lineWidth = 2;
        ctx.beginPath();
        ctx.moveTo(player.x, player.y);
        ctx.lineTo(
            player.x + Math.cos(player.direction) * config.playerSize * 1.5,
            player.y + Math.sin(player.direction) * config.playerSize * 1.5
        );
        ctx.stroke();
    }

    // 绘制子弹
    ctx.fillStyle = '#ffff00';
    for (const bulletId in gameState.bullets) {
        const bullet = gameState.bullets[bulletId];
        ctx.beginPath();
        ctx.arc(bullet.x, bullet.y, config.bulletSize, 0, Math.PI * 2);
        ctx.fill();
    }
}

// 处理键盘输入
window.addEventListener('keydown', (e) => {
    switch (e.key.toLowerCase()) {
        case 'w': gameState.inputState.up = true; break;
        case 's': gameState.inputState.down = true; break;
        case 'a': gameState.inputState.left = true; break;
        case 'd': gameState.inputState.right = true; break;
    }
});

window.addEventListener('keyup', (e) => {
    switch (e.key.toLowerCase()) {
        case 'w': gameState.inputState.up = false; break;
        case 's': gameState.inputState.down = false; break;
        case 'a': gameState.inputState.left = false; break;
        case 'd': gameState.inputState.right = false; break;
    }
});

// 处理鼠标移动
canvas.addEventListener('mousemove', (e) => {
    const rect = canvas.getBoundingClientRect();
    gameState.inputState.mouseX = e.clientX - rect.left;
    gameState.inputState.mouseY = e.clientY - rect.top;
});

// 处理鼠标点击
canvas.addEventListener('mousedown', (e) => {
    if (e.button === 0) { // 左键
        gameState.inputState.shooting = true;
    }
});

// 游戏循环
let lastTime = 0;
function gameLoop(timestamp) {
    const deltaTime = timestamp - lastTime;
    lastTime = timestamp;

    // 每秒发送30次输入 (约33ms间隔)
    if (timestamp - gameState.lastUpdateTime > 33) {
        sendInputState();
        gameState.lastUpdateTime = timestamp;
    }

    requestAnimationFrame(gameLoop);
}

// 启动游戏循环
requestAnimationFrame(gameLoop);