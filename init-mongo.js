db.createUser(
    {
        user: "bryan",
        pwd: "password",
        roles: [
            {
                role: "readWrite",
                db: "socially"
            }
        ]
    }
)
