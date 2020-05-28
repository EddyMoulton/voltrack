export class StockChangeViewModel {
  public value: number;
  public simpleReturn: number;
  public annualReturn: number;
  public scaledAnnualReturn: number | undefined;

  constructor(value?: number, simpleReturn?: number, annualReturn?: number, scaledAnnualReturn?: number) {
    this.value = value || 0;
    this.simpleReturn = simpleReturn || 0;
    this.annualReturn = annualReturn || 0;
    this.scaledAnnualReturn = scaledAnnualReturn;
  }

  generate(currentValue: number, baseValue: number, yearsHeld: number, scaleBase?: number, returnBaseValue?: number) {
    if (!returnBaseValue) {
      returnBaseValue = baseValue;
    }

    this.value = currentValue - baseValue;
    this.simpleReturn = 100 * this.value / returnBaseValue;

    this.annualReturn = Math.pow((this.simpleReturn + 1), (1 / yearsHeld)) - 1;

    if (scaleBase) {
      this.scaledAnnualReturn = 100 * this.annualReturn / scaleBase;
    }
  }

  sumarise(stockChanges: StockChangeViewModel[]) {
    this.value = stockChanges.map((change) => change.value).reduce((total, value) => total + value);
    this.scaledAnnualReturn = stockChanges.map((change) => change.scaledAnnualReturn).reduce((total, value) => (total || 0) + (value || 0));
  }
}