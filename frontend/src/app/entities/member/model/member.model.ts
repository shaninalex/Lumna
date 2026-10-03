export interface Member {
    id: number;
    workspaceId: number;
    email: string;
    fullName: string;
    image?: string;
    role: string;
    dateJoined: Date;
}
