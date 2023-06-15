import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation.model';
import { AccommodationService } from 'src/app/service/accommodation.service';

@Component({
  selector: 'app-show-accommodations',
  templateUrl: './show-accommodations.component.html',
  styleUrls: ['./show-accommodations.component.css']
})
export class ShowAccommodationsComponent implements OnInit {

  public dataSource = new MatTableDataSource<Accommodation>();
  public displayedColumns = ['Name', 'Benefits', 'Address'];
  public accommodations: Accommodation[] = [];
  public accommodation: Accommodation | undefined = undefined;


  constructor(
    private router: Router,
    private accommodationService : AccommodationService
  ) { }

  ngOnInit(): void {

    this.accommodationService.GetAllAccommodations().subscribe((data: any) => {
      this.accommodations = data.accommodations;
      console.log(this.accommodations);
      this.dataSource.data = this.accommodations;
      console.log(this.dataSource.data);
    })
  }

}
