import {Address} from "./address.model";
import { Availability } from "./availability.model";

export class Accommodation {
    constructor(
        public id: string = '',
        public name: string = '',
        public minNumberOfGuests: number = 0,
        public maxNumberOfGuests: number = 0,
        public hostId: string = '',
        public benefits: string[] = [],
        public address: Address = new Address(),
        public startDate: string = '',
        public endDate: string = '',
        public price: number = 0,
        public totalPrice: number  = 0,
        public priceSelection: string = '',
        public availabilities: Availability[] = [],
        public isSuperHost: boolean = false, 

      ) {}
}
