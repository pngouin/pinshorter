export class Link {
    id: string;
    title: string;
    url: string;
    count: number;
    api_point: string;
    createdAt: Date;
    user : {
        id: string;
        name: string;
    }
}
