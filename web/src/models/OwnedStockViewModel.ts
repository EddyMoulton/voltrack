export class OwnedStockViewModel {
  public code: string = '';
  public quantity: number = 0;
  public currentValue: number = 0;
  public paidValue: number = 0;
  public difference: number = 0;
  public purchaseDate: Date = new Date();
  public totalDividends: number = 0;
}