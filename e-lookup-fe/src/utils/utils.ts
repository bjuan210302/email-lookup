const ELOOKUP_BACKEND_QUERY_URL = import.meta.env.VITE_ELOOKUP_BACKEND_QUERY_URL

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

export const makeQueryRequest = async (term: string, page: number, maxPerPage: number) => {
  const query = ELOOKUP_BACKEND_QUERY_URL + new URLSearchParams(
    {
      'word': term,
      'page': String(page),
      'max_per_page': String(maxPerPage)
    }
  )
  const res = await fetch(query, {
    method: 'get',
    headers: new Headers({
      'Authorization': 'Basic ' + 'username:password',
      'Content-Type': 'application/json'
    }),
  })
  return res.json() as Promise<QueryHits>
}