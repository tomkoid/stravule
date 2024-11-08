import { PUBLIC_BACKEND_URL } from '$env/static/public';

export interface Filters {
  include: string[]
  exclude: string[]
}

// returns all filters for the current user 
export async function loadFilters(sid: string, canteen: string): Promise<Filters> {
  let params = new URLSearchParams()

  params.append("sid", sid)
  params.append("canteen", canteen)
  let req = await fetch(`${PUBLIC_BACKEND_URL}/api/v1/filters?${params.toString()}`, {
    method: "GET",
  })

  if (req.status != 200) {
    throw new Error(await req.text())
  }

  let data: Filters = await req.json()

  return data
}

export async function removeFilter(sid: string, canteen: string, filterString: string) {
  let params = new URLSearchParams()

  params.append("sid", sid)
  params.append("canteen", canteen)
  let req = await fetch(`${PUBLIC_BACKEND_URL}/api/v1/filters?${params.toString()}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      category: "",
      value: filterString,
    }),
  })

  if (req.status != 200) {
    throw new Error(await req.text())
  }

  let data: Filters = await req.json()

  return data
}
