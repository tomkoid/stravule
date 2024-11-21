import { env } from '$env/dynamic/public';
import { errors } from '$lib/stores/errors.svelte';
import { pageLoading } from '$lib/stores/page.svelte';

export interface Order {
  datum: string,
  druh: string,
  chod: string,
  pocet: number,
  cena: string,
  omezeni: string,
  polevka: string,
  id: number,
  nazev: string,
  druhPopis: string,
  druhChod: string,
  databaze: string,
  delsiPopis: string,
  alergenyText: string,
  zkratkaProduktu: string,
  multipleNazev: string,
  version: number,
  casKonec: string,
  casOdhlaseni: string,
}
// returns if logged in
export async function loadOrders(sid: string, canteen: string, pickOrders: boolean): Promise<Order[][]> {
  pageLoading.value = true;

  let params = new URLSearchParams()

  params.append("sid", sid)
  params.append("canteen", canteen)
  params.append("user_hash", localStorage.getItem("user_hash")!)
  params.append("pick", Boolean(pickOrders).toString())
  let req = await fetch(`${env.PUBLIC_BACKEND_URL}/api/v1/orders?${params.toString()}`, {
    method: "GET",
  })

  if (req.status != 200) {
    errors.add(await req.text())
    throw new Error(await req.text())
  }

  let data: Order[][] = await req.json()

  pageLoading.value = false;

  return data
}

// confirms the selection by stravule and set orders 
export async function sendOrders(sid: string, canteen: string) {
  pageLoading.value = true;

  let params = new URLSearchParams()

  params.append("sid", sid)
  params.append("canteen", canteen)
  params.append("user_hash", localStorage.getItem("user_hash")!)
  let req = await fetch(`${env.PUBLIC_BACKEND_URL}/api/v1/save_orders?${params.toString()}`, {
    method: "POST",
  })

  if (req.status != 200) {
    errors.add(await req.text())
    pageLoading.value = false;
    throw new Error(await req.text())
  }

  pageLoading.value = false;
}
