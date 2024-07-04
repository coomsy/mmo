
export const getPagingQuery = (page: number, limit: number): string =>
    `limit=${limit}&offset=${page ? page * limit : 0}`;
