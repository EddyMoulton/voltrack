import { StockChangeViewModel } from './StockChangeViewModel';

export class TransactionSummaryViewModel {
  public code: string;
  public purchaseDate: Date;
  public quantity: number;
  public capital: StockChangeViewModel;
  public dividends: StockChangeViewModel;
  public total: StockChangeViewModel;

  constructor(code: string, quantity: number, costValue: number, currentValue: number, dividendValue: number, purchaseDate: Date, totalCostValue?: number) {
    this.code = code;
    this.quantity = quantity;
    this.purchaseDate = purchaseDate;

    const now = new Date();
    const yearsHeld = (now.getTime() - purchaseDate.getTime()) / (365 * 24 * 60 * 60 * 1000);


    this.capital = new StockChangeViewModel(currentValue, costValue, yearsHeld, totalCostValue);
    this.dividends = new StockChangeViewModel(dividendValue, 0, yearsHeld, totalCostValue, costValue);
    this.total = new StockChangeViewModel(currentValue + dividendValue, costValue, yearsHeld, totalCostValue);
  }
}