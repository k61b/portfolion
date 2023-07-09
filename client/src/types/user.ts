export interface IUser {
  username: string
  bookmarks: {
    symbol: string
    added_price: number
    pices: number
  }
  value: number
  profit_and_loss: number
}