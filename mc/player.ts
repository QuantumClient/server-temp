import {ServerClient} from "minecraft-protocol";
import {db} from "./db";
import {QueryError, RowDataPacket} from "mysql2";
export class Player {

    private client: ServerClient


    constructor(client: ServerClient) {
        this.client = client;
    }

    getKey(callback: (result: string) => any) {
        const queryString = 'SELECT `keys`, `player_username` FROM `verified_mc_accounts` WHERE `player_uuid` =?'
        db.query(queryString, this.client.uuid, (err, result) => {
            if (err) console.log(err)
            const row = (<RowDataPacket> result)[0];

            if (row) {
                return callback(row.keys as string)
            } else {
                return callback(this.insertNew())
            }
        });
    }

    private insertNew(): string {
        const insetString = "INSERT INTO verified_mc_accounts (`keys`, `player_uuid`, `player_username`) VALUES (?,?,?)"
        const keyt = Player.genKey()
        db.query(insetString, [keyt, this.client.uuid, this.client.username], (err, result) => {
            if (err) console.log(err)

            return keyt
        })
        return keyt
    }

    private static genKey(): string {
       return  Math.random().toString(36).substr(2, 9);
    }

    public static formatKey(key: string): string {
        return key.match(/.{1,3}/g).join('-');
    }

}
