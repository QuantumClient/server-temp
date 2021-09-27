import {createServer, Server, ServerOptions} from "minecraft-protocol";
import Jimp from 'jimp';
import {Player} from "./player";
import {db} from "./db";

const motd = '\u00a75\u00a7nQuantum\u00a79 Account Verification Server'
const img = Jimp.read("https://quantumclient.org/favicon.png")
        .then(img => img.resize(64, 64))
        .then(async img => {
            return await img.getBase64Async(Jimp.MIME_PNG);
        });

const serverConfig: ServerOptions = {
    'online-mode': true,
     beforePing: async (response, client, answerToPing) => {
        const pingResponse = {
            version: {
                name: '1.8-1.17',
                protocol: client.protocolVersion,
            },
            players: {
                online: 420,
                max: 69,
                sample: [],
            },
            description: { text: motd },
            favicon: await img

        };
        client.write('server_info', { response: JSON.stringify(pingResponse) });
    }
}



const server: Server = createServer(serverConfig);

server.on('login', (client) => {


    const p = new Player(client)
    p.getKey((token) => {
        client.end( `§5Account Verified \n Your code is \n §d§l${Player.formatKey(token)}`)
    })
});
