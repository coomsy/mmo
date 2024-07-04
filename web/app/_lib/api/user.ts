import { User } from "../types/user";
import { API_URL, API_POST_HEADERS } from "../util/constant";
import { getPagingQuery } from "../util/queryBuilder";

const UserAPI = {
    all: (search: string, page: number, limit: number = 10) =>
        fetch(`${API_URL}/guild?${getPagingQuery(page, limit)}`),

    byId: (id: string) =>
        fetch(`${API_URL}/guild/${id}`),

    create: async (guild: User, token: string): Promise<User | null>  => 
        await fetch(`${API_URL}/guild`, {
                method: 'POST',
                headers: {...API_POST_HEADERS, ...{
                        'Authorization': token
                    }
                },
                body: JSON.stringify({guild})
            }
        ).then(resp => {
            if (!resp.ok) {
                console.error('_lib.api.GuildAPI.create [%s]', resp.statusText)
                return null
            }
            return resp.json() as Promise<User>
        }),

}