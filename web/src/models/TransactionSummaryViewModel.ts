import { StockChangeViewModel } from './StockChangeViewModel';

export class TransactionSummaryViewModel {
  public code: string;
  public purchaseDate: Date | undefined;
  public quantity: number;
  public capital: StockChangeViewModel;
  public dividends: StockChangeViewModel;
  public total: StockChangeViewModel;

  constructor(
    code?: string,
    purchaseDate?: Date,
    quantity?: number,
    captial?: StockChangeViewModel,
    dividends?: StockChangeViewModel,
    total?: StockChangeViewModel,
  ) {
    this.code = code || '';
    this.purchaseDate = purchaseDate;
    this.quantity = quantity || 0;
    this.capital = captial || new StockChangeViewModel();
    this.dividends = dividends || new StockChangeViewModel();
    this.total = total || new StockChangeViewModel();
  }

  generate(
    code: string,
    quantity: number,
    costValue: number,
    currentValue: number,
    dividendValue: number,
    purchaseDate: Date,
    totalCostValue?: number,
  ) {
    this.code = code;
    this.quantity = quantity;
    this.purchaseDate = purchaseDate;

    const now = new Date();
    const yearsHeld =
      (now.getTime() - purchaseDate.getTime()) / (365 * 24 * 60 * 60 * 1000);

    this.capital.generate(currentValue, costValue, yearsHeld, totalCostValue);
    this.dividends.generate(
      dividendValue,
      0,
      yearsHeld,
      totalCostValue,
      costValue,
    );
    this.total.generate(
      currentValue + dividendValue,
      costValue,
      yearsHeld,
      totalCostValue,
    );
  }

  summarise(transactions: TransactionSummaryViewModel[]) {
    this.code = 'Total';
    this.capital = new StockChangeViewModel();
    this.capital.sumarise(
      transactions.map((t: TransactionSummaryViewModel) => t.capital),
    );
    this.dividends = new StockChangeViewModel();
    this.dividends.sumarise(
      transactions.map((t: TransactionSummaryViewModel) => t.dividends),
    );
    this.total = new StockChangeViewModel();
    this.total.sumarise(
      transactions.map((t: TransactionSummaryViewModel) => t.total),
    );
  }
}
