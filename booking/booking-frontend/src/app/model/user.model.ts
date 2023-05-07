import {Address} from "./address.model";

export class User {
  constructor(
    public id: string = '',
    public firstName: string = '',
    public lastName: string = '',
    public email: string = '',
    public password: string = '',
    public role: string = '',
    public address: Address = new Address()
  ) {}
}

export class LoginRequest {
  constructor(public email: string = '', public password: string = '') {}
}
