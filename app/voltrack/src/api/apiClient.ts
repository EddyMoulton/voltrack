import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';

export class ApiClient {
  private readonly getOwnedStockLogsUrl = "/stocks/history";
  private readonly GenerateSummaryLogsUrl = "/reporting/generate";
  private readonly GetAllTransactionsUrl = "/stocks/transactions";
  private readonly AddTransactionUrl = "/stocks/transactions";
  private readonly GetAllStocksUrl = "/stocks";

  private axios: AxiosInstance;

  constructor(axiosConfig?: AxiosRequestConfig | undefined) {
    this.axios = axios.create(axiosConfig);
  }

  public getOwnedStockLogs() {
    return this.axios.get(this.getOwnedStockLogsUrl)
  }

  public GenerateSummaryLogs() {
    return this.axios.get(this.GenerateSummaryLogsUrl)
  }

  public GetAllTransactions() {
    return this.axios.get(this.GetAllTransactionsUrl)
  }

  public AddTransaction() {
    return this.axios.get(this.AddTransactionUrl)
  }

  public GetAllStocks() {
    return this.axios.get(this.GetAllStocksUrl)
  }
}