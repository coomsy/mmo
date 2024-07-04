import { Guild } from "../types/guild";
import { API_URL, API_POST_HEADERS } from "../util/constant";
import { getPagingQuery } from "../util/queryBuilder";



const GuildAPI = {
    all: (search: string, page: number, limit: number = 10) =>
        fetch(`${API_URL}/guild?${getPagingQuery(page, limit)}`
            ).then(resp => { return resp.json() as Promise<[Guild]> } ),

    byId: async (id: string): Promise<Guild | null>  => 
        fetch(`${API_URL}/guild/${id}`
            ).then(resp => { 
                if (!resp.ok) {
                    console.error('_lib.api.GuildAPI.byId [%s]', resp.statusText)
                    return null
                }
                return resp.json() as Promise<Guild>
             } ),

    create: async (guild: Guild, token: string): Promise<Guild | null>  => 
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
            return resp.json() as Promise<Guild>
        }),

}


export default GuildAPI;