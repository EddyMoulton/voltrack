import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import { OwnedStockViewModel } from '@/models/OwnedStockViewModel';
import { AddTransactionDto } from './models/AddTransactionDto';

export class ApiClient {
  private readonly getOwnedStockLogsUrl = "/stocks/history";
  private readonly GenerateSummaryLogsUrl = "/reporting/generate";
  private readonly GetAllTransactionsUrl = "/stocks/transactions";
  private readonly AddTransactionUrl = "/stocks/transactions";
  private readonly GetAllStocksUrl = "/stocks";
  private readonly GetCurrentStocksUrl = "/stocks/current";

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
    return this.axios.get(this.baseUrl + this.getOwnedStockLogsUrl)
  }

  public GenerateSummaryLogs() {
    return this.axios.get(this.baseUrl + this.GenerateSummaryLogsUrl)
  }

  public GetAllTransactions() {
    return this.axios.get(this.baseUrl + this.GetAllTransactionsUrl)
  }

  public AddTransaction(dto: AddTransactionDto) {
    return this.axios.post(this.baseUrl + this.AddTransactionUrl, dto)
  }

  public GetAllStocks() {
    return this.axios.get(this.baseUrl + this.GetAllStocksUrl)
  }

  public GetCurrentStocks(): Promise<OwnedStockViewModel[]> {
    return this.axios.get(this.baseUrl + this.GetCurrentStocksUrl).then((result) => {
      if (result.status == 200) {
        return result.data.currentStocks.map((data: OwnedStockViewModel) => {
          const entry = data as OwnedStockViewModel;

          entry.currentValue = entry.currentValue / 10000;
          entry.totalValue = entry.totalValue / 10000;
          entry.paidValue = entry.paidValue / 10000;
          entry.difference = entry.difference / 10000;

          return entry;
        });
      } else {
        throw Error(result.statusText);
      }
    })
  }
}