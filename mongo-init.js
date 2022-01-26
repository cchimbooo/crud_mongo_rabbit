const db = new Mongo().getDB('teste');

//db.createUser(
//    {
//        user: "user",
//        pwd: "pwd",
//        roles: [
//            {
//                role: "readWrite",
//                db: "teste"
//            }
//        ]
//    }
//);


db.getCollection('pessoa').createIndex({nome: 1}, {unique: true});
db.getCollection('teste').createIndex({nome: 1}, {unique: true});