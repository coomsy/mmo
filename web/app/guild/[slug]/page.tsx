import GuildAPI from '@/app/_lib/api/guild'
import { Suspense } from 'react'

// https://nextjs.org/docs/app/building-your-application/data-fetching/patterns
export default async function Page({params: { slug }}: { params: { slug: string }}) {
  const guild = await GuildAPI.byId(slug).then(r => {return r})

  return (
    <>
      <h1>Guilds</h1>
      <Suspense fallback={<div>Loading...</div>}>
        <ul>
          { guild != null ?
            <li key={guild.Id}>{guild.Id} | {guild.Name} | {guild.CreatedAt}</li>
            : 
            <p>No existo</p> 
          }
        </ul>
      </Suspense>
    </>
  )
}