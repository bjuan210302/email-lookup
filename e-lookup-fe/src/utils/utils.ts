const ELOOKUP_HOST = import.meta.env.VITE_ELOOKUPBE_HOST + "api/v1"

export type Email = {
  _id: string,
  date: string,
  from: string,
  to: string[],
  subject: string,
  content: string,
  highlight: string[]
}

export type QueryHits = {
  totalHits: number,
  hits: Email[]
}

export type SearchConfig = {
  zIndex: string;
  resultsPerPage: number;
}

export const makeQueryRequest = async (term: string, page: number, config: SearchConfig) => {
  const { zIndex, resultsPerPage } = config
  const query = ELOOKUP_HOST + "/lookup?" + new URLSearchParams(
    {
      'word': term,
      'page': String(page),
      'max_per_page': String(resultsPerPage),
      'index_name': String(zIndex)
    }
  )

  const res = await fetch(query, {
    headers: new Headers({
      'Authorization': sessionStorage.getItem('auth') as string,
      'Content-Type': 'application/json'
    }),
  })

  const body = await res.json()
  if (res.status > 200) {
    throw new Error(body)
  } 

  const { hits, totalHits } = body as QueryHits
  let numPages = Math.trunc(totalHits / resultsPerPage)
  if (totalHits % resultsPerPage !== 0) {
    numPages++
  }
  
  return { hits, totalHits, numPages }
}

export const getIndexes = async (auth?: string) => {
  const link = ELOOKUP_HOST + "/indexes"
  auth = auth || sessionStorage.getItem('auth') as string

  const res = await fetch(link, {
    headers: new Headers({
      'Authorization': auth,
      'Content-Type': 'application/json'
    }),
  })

  const body = await res.json()
  if (res.status > 200) {
    throw new Error(body)
  } 

  return body as string[]
}

export const zAuthenticate = async (zUser: string, zPass: string) => {
  const auth = `Basic ${btoa(`${zUser}:${zPass}`)}`

  // Using this to verify credentials
  // If this doesn't raise an error then set the auth
  const indexes = await getIndexes(auth);
  sessionStorage.setItem('auth', auth)
  sessionStorage.setItem('zUser', zUser)
  sessionStorage.setItem('zPass', zPass)

  return indexes
}
