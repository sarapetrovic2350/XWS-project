import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import { Observable } from 'rxjs';
import { Accommodation } from '../model/accommodation.model';

@Injectable({
  providedIn: 'root'
})
export class AccommodationService {

  apiHost: string = 'http://localhost:8080/api/accommodation/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Access-Control-Allow-Origin' : '*',
  'Access-Control-Allow-Methods' : 'GET,HEAD,OPTIONS,POST,PUT', 'Access-Control-Allow-Headers' : 'Origin, X-Requested-With, Content-Type, Accept, x-client-key, x-client-token, x-client-secret, Authorization' }  );

  constructor(private http: HttpClient) { }

  createAccommodation(accommodation: Accommodation) {
    console.log(accommodation);
    return this.http.post<Accommodation>(this.apiHost , accommodation);
  }

  searchAccommodations(searchAccommodations: any) {
    console.log(searchAccommodations);
    return this.http.post<Accommodation[]>(this.apiHost + 'search' , searchAccommodations);
  }
}
