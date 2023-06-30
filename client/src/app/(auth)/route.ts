import { postFetcher } from "@utils/fetch"

export async function POST(path: string, body: any) {
  return await postFetcher(path, body)
}
