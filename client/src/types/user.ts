export interface IUser {
  avatar: string
  username: string
  bookmarks: Array<Bookmark>
  value: number
  profit_and_loss: number
}

interface Bookmark {
  symbol: string
  added_price: number
  pices: number
}