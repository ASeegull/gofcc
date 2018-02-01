import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-testresult',
  templateUrl: './testresult.component.html',
  styleUrls: ['./testresult.component.css']
})
export class TestresultComponent implements OnInit {

  @Input('results') results;
  constructor() { }

  ngOnInit() {
  }

}
