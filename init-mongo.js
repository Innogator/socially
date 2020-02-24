db.createUser(
    {
        user: "bryan",
        pwd: "gui",
        roles: [
            {
                role: "readWrite",
                db: "socially"
            }
        ]
    }
)
