import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import { OwnedStockViewModel } from '@/models/OwnedStockViewModel';
import { AddTransactionDto } from './models/AddTransactionDto';
import { TransactionSummaryViewModel } from '@/models/TransactionSummaryViewModel';
import { TransactionSummaryDto } from '@/models/TransactionSummaryDto';

export class ApiClient {
  private readonly getOwnedStockLogsUrl = "/stocks/history";
  private readonly generateSummaryLogsUrl = "/reporting/generate";
  private readonly getAllTransactionsUrl = "/stocks/transactions";
  private readonly addTransactionUrl = "/stocks/transactions";
  private readonly getAllStocksUrl = "/stocks";
  private readonly getCurrentStocksUrl = "/stocks/current";
  private readonly getTransactionSummariesUrl = "/transactions/summaries"

  private axios: AxiosInstance;
  private baseUrl: string;

  constructor(baseUrl: string, axiosConfig?: AxiosRequestConfig | undefined) {
    this.baseUrl = baseUrl.replace(/\/$/, ""); // Ensure no trailing slash

    if (!axiosConfig) {
      axiosConfig = {};
    }

    axiosConfig!.headers = { ...axiosConfig?.headers, ...{ "Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "access-control-allow-origin, access-control-allow-headers", "Content-Type": "text/plain" } }

    this.axios = axios.create(axiosConfig);
  }

  public getOwnedStockLogs() {
    return this.axios.get(this.baseUrl + this.getOwnedStockLogsUrl);
  }

  public GenerateSummaryLogs() {
    return this.axios.get(this.baseUrl + this.generateSummaryLogsUrl);
  }

  public GetAllTransactions() {
    return this.axios.get(this.baseUrl + this.getAllTransactionsUrl);
  }

  public AddTransaction(dto: AddTransactionDto) {
    return this.axios.post(this.baseUrl + this.addTransactionUrl, dto);
  }

  public GetAllStocks() {
    return this.axios.get(this.baseUrl + this.getAllStocksUrl);
  }

  public GetCurrentStocks(): Promise<OwnedStockViewModel[]> {
    return this.axios.get(this.baseUrl + this.getCurrentStocksUrl).then((result) => {
      if (result.status == 200) {
        return result.data.currentStocks.map((data: OwnedStockViewModel) => {
          const entry = data as OwnedStockViewModel;

          entry.currentValue = entry.currentValue / 10000;
          entry.paidValue = entry.paidValue / 10000;
          entry.difference = entry.difference / 10000;
          entry.totalDividends = entry.totalDividends / 10000;

          return entry;
        });
      } else {
        throw Error(result.statusText);
      }
    })
  }

  public async GetTransactionSummaries(): Promise<TransactionSummaryViewModel[]> {
    const result = await this.axios.get(this.baseUrl + this.getTransactionSummariesUrl)

    console.log(result.data.transactions)

    const totalCost = result.data.transactions.reduce((total: number, t: TransactionSummaryDto) => {
      return total + t.cost! || 0;
    });

    const a = result.data.transactions.map((t: TransactionSummaryDto) => {
      return new TransactionSummaryViewModel(
        t.code!,
        t.quantity!,
        t.cost! * t.quantity!,
        t.value! * t.quantity!,
        t.dividendValue!,
        new Date(t.date!),
        totalCost
      );
    });
    console.log(a)
    return a
  }
}