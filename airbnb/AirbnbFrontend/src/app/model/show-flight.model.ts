export class ShowFlight {
  constructor(
    public id: string = '',
    public departureDateTime: string = '',
    public arrivalDateTime: string = '',
    public departure: string = '',
    public arrival: string = '',
    public price: number = 0,
    public totalNumberOfSeats: number = 0,
    public availableSeats: number = 0
  ) {}
}
