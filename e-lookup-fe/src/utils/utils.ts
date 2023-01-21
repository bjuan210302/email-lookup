const ELOOKUP_BACKEND_QUERY_URL = import.meta.env.VITE_ELOOKUPBE_HOST + "api/v1/lookup?"

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
  zUser: string;
  zPass: string;
  zIndex: string;
  resultsPerPage: number;
}

export const makeQueryRequest = async (term: string, page: number, config: SearchConfig) => {
  const { zUser, zPass, zIndex, resultsPerPage } = config
  const query = ELOOKUP_BACKEND_QUERY_URL + new URLSearchParams(
    {
      'word': term,
      'page': String(page),
      'max_per_page': String(resultsPerPage),
      'index_name': String(zIndex)
    }
  )
  const { hits, totalHits } = await fetch(query, {
    method: 'get',
    headers: new Headers({
      'Authorization': `Basic ${zUser}:${zPass}`,
      'Content-Type': 'application/json'
    }),
  }).then(r => r.json() as Promise<QueryHits>)

  let numPages = Math.trunc(totalHits / resultsPerPage)
  if (totalHits % resultsPerPage !== 0) {
    numPages++
  }

  return { hits, totalHits, numPages }
}