const mineflayer = require('mineflayer');
const {pathfinder, Movements, goals } = require('mineflayer-pathfinder');
const GoalFollow = goals.GoalFollow

let target = null

const bot = mineflayer.createBot({
    host: 'localhost',
    port: '53456',
    username: 'ClaudinhoBot'
})

function onCheating(username, message){
    switch(message){
        case 'noite':
            bot.chat('/time set night')
            break;
        case 'dia':
            bot.chat('/time set day')
            break;
        case 'limpe os inventários':
            bot.chat('/clear @a')
            break;
        case 'pare a chuva':
            bot.chat('/weather clear')
            break;
        case 'faça chuver': 
            bot.chat('/weather rain')
            break;
    }
}

function onTalking(username, message){
    switch(message){
        case 'Bom dia!':
            bot.chat('Bom dia!!! Como está você?')
            break;
        case 'Boa noite!':
            bot.chat('Bom noite!!! A lua está linda, não é?')
            break;
        case 'Boa tarde!':
            bot.chat('Boa tarde!!! Que calor da mizera')
            break;
        case 'Como vai?' || 'como vai?':
            bot.chat('Vou bem! Obrigado por perguntar. E você?')
            break;
        case 'Estou meio mal':
            bot.chat('Sério?? :( oq aconteceu???')
            break;
        case 'Estou muito bem!!':
            bot.chat('Que ótimo, fico muito feliz')
            break;
    } 
}

function onStating(username, message){
    switch(message){
        case 'corra':
            bot.setControlState('sprint', true)
            break;
        case 'pare':
            bot.clearControlStates()
            break;
        case 'frente':
            bot.setControlState('forward', true)
            break;
        case 'direita':
            bot.setControlState('right', true)
            break;
        case 'esquerda':
            bot.setControlState('left', true)
            break;
        case 'atras':
            bot.setControlState('back', true)
            break;
        case 'pule':
            bot.setControlState('jump', true)
            bot.setControlState('jump', false)
            break;
        case 'pule até cansar':
            bot.setControlState('jump', true)
            break;
        case 'olhe pra mim': 
            bot.lookAt(username)
            break;
        case 'ataque':
            entity = bot.nearestEntity()

            if(entity){
                bot.attack(entity, true)
            } else{
                bot.chat('Não tem nada pra eu matar')
            }
            break;
        case 'desmonte':
            bot.dismount()
            break;
        case 'tpy':
            bot.entity.position.y += 10
            break;
        case 'tpx':
            bot.entity.position.x += 10
            break;
        case 'posicao':
            bot.chat(bot.entity.position.toString())
            break;
        case 'yp':
            bot.chat(`Yaw ${bot.entity.yaw}, pitch: ${bot.entity.pitch}`)
            break
        case 'me':
            hitPlayer()
    }
}

function lookAtPlayer(){
    playerEntity = bot.nearestEntity()
    if(!playerEntity) return
    pos = playerEntity.position.offset(0, playerEntity.height, 0)
    bot.lookAt(pos)
}

// function followPlayer(){
//     const playerI = bot.players['pythonul']
//     if(!player || !player.entity){
//         return
//     }
//     const mineData = require('minecraft-data')(bot.version)
//     const movements = new Movements(bot, mineData)

//     bot.pathfinder.setMovements(movements)

//     const goal = new GoalFollow(playerI.entity, 2)
//     bot.pathfinder.setGoal(goal, true)
// }

function hitPlayer(){
    playerEntity = bot.nearestEntity()
    if(!playerEntity) return
    pos = playerEntity.position.offset(0, playerEntity.height, 0)
    bot.attack(playerEntity, true)
}

bot.on('login', () => {
    bot.chat('Salve rapeizi')
})

bot.on('chat', (username, message) => {
    onCheating(username, message)
    onTalking(username, message)
    onStating(username, message)
})

bot.on('physicTick', lookAtPlayer)

bot.on('rain', () => {
    bot.chat('/weather clear')
})