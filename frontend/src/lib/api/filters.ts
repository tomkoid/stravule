import { PUBLIC_BACKEND_URL } from '$env/static/public';
import { errors } from '$lib/stores/errors.svelte';

export interface Filters {
  include: string[]
  exclude: string[]
}

// returns all filters for the current user 
export async function loadFilters(): Promise<Filters> {
  let params = new URLSearchParams()

  params.append("user_hash", localStorage.getItem("user_hash")!)
  let req = await fetch(`${PUBLIC_BACKEND_URL}/api/v1/filters?${params.toString()}`, {
    method: "GET",
  })

  if (req.status != 200) {
    errors.add(await req.text())
    throw new Error(await req.text())
  }

  let data: Filters = await req.json()

  return data
}

export async function addFilter(filterString: string, category: string): Promise<Filters> {
  let params = new URLSearchParams()

  params.append("user_hash", localStorage.getItem("user_hash")!)
  let req = await fetch(`${PUBLIC_BACKEND_URL}/api/v1/filters?${params.toString()}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      category: category,
      value: filterString,
    }),
  })

  if (req.status != 200) {
    errors.add(await req.text())
    throw new Error(await req.text())
  }

  let data: Filters = await req.json()

  return data
}

export async function removeFilter(filterString: string): Promise<Filters> {
  let params = new URLSearchParams()

  params.append("user_hash", localStorage.getItem("user_hash")!)
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
    errors.add(await req.text())
    throw new Error(await req.text())
  }

  let data: Filters = await req.json()

  return data
}
