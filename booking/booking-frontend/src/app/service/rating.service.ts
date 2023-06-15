import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { RatingHost } from '../model/rating-host.model';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RatingService {

  apiHost: string = 'http://localhost:8000/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Access-Control-Allow-Origin' : '*',
    'Access-Control-Allow-Methods' : 'GET,HEAD,OPTIONS,POST,PUT', 'Access-Control-Allow-Headers' : 'Origin, X-Requested-With, Content-Type, Accept, x-client-key, x-client-token, x-client-secret, Authorization' }  );
  headers2: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient) { }

  createRatingForHost(ratingHost: RatingHost): Observable<any> {
    return this.http.post<RatingHost>(this.apiHost + 'createRatingForHost', ratingHost, {headers: this.headers2});
  }

}
