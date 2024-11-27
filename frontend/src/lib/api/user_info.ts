import { PUBLIC_BACKEND_URL } from '$env/static/public';
import { errors } from '$lib/stores/errors.svelte';
import { pageLoading } from '$lib/stores/page.svelte';

export interface UserInfo {
  konto: string;
  vs: string;
  jmeno: string;
  email: string;
  vydej: number;
  limit: string;
  vicenasobny: boolean;
  pocetJidel: number;
  podrobnosti: number;
  editaceJidelnicku: number;
  id: string;
  cislo: string;
  bakalarId: string;
  prihlaska: number;
  pasivni: boolean;
  verze: string;
  text: string;
  mena: string;
  dochazka: string;
  nazevJidelny: string;
}

// loads user info into the store 
export async function getUserInfo(sid: string, canteen: string): Promise<UserInfo> {
  pageLoading.value = true;

  let params = new URLSearchParams()

  params.append("sid", sid)
  params.append("canteen", canteen)
  params.append("user_hash", localStorage.getItem("user_hash")!)
  let req = await fetch(`${PUBLIC_BACKEND_URL}/api/v1/account_info?${params.toString()}`, {
    method: "GET",
  })

  if (req.status != 200) {
    errors.add(await req.text())
    throw new Error(await req.text())
  }

  let data: UserInfo = await req.json()

  localStorage.setItem("jmeno", data.jmeno)

  pageLoading.value = false;

  return data
}
