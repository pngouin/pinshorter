export class Connection {
    name: string;
    password: string;
}

export class User {
    id: string;
    name: string;
    token?: string
}

export class Token {
    token: string;
}