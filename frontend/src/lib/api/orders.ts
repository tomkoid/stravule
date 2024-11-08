import { PUBLIC_BACKEND_URL } from '$env/static/public';
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
  params.append("pick", Boolean(pickOrders).toString())
  let req = await fetch(`${PUBLIC_BACKEND_URL}/api/v1/orders?${params.toString()}`, {
    method: "GET",
  })

  if (req.status != 200) {
    throw new Error(await req.text())
  }

  let data: Order[][] = await req.json()

  pageLoading.value = false;

  return data
}
