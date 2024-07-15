db.log.insertOne({"message": "Database created."});
db.createUser(
    {
        user: _getEnv("MONGO_INITDB_USERNAME"),
        pwd: _getEnv("MONGO_INITDB_PASSWORD"),
        roles: [
            {role: "readWrite", db: _getEnv("MONGO_INITDB_DATABASE")}
        ]
    }
);