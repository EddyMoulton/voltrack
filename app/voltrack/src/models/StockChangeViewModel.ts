export class StockChangeViewModel {
  public value: number;
  public simpleReturn: number;
  public annualReturn: number;
  public scaledAnnualReturn: number | undefined;

  constructor(currentValue: number, baseValue: number, yearsHeld: number, scaleBase?: number, returnBaseValue?: number) {
    if (!returnBaseValue) {
      returnBaseValue = baseValue;
    }

    this.value = currentValue - baseValue;
    this.simpleReturn = this.value / returnBaseValue;

    this.annualReturn = Math.pow((this.simpleReturn + 1), (1 / yearsHeld)) - 1;

    if (scaleBase) {
      this.scaledAnnualReturn = this.annualReturn / scaleBase;
    }
  }
}