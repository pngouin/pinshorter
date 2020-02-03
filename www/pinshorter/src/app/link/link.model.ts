export class Link {
    id: string;
    title: string;
    url: string;
    count: number;
    apiPoint: string;
    createdAt: Date;
    user : {
        id: string;
        name: string;
    }
}
