import { Component, ElementRef, ViewChild } from '@angular/core';
import { MainService } from '../../services/main.service';

@Component({
  selector: 'app-editor',
  templateUrl: './editor.component.html',
  styleUrls: ['./editor.component.css']
})
export class EditorComponent {
  initialcode: string;
  solution: string;
  @ViewChild('codeinput') public codeinput: ElementRef;

  constructor(public mainService: MainService) { }

  public focus(): void {
    this.codeinput.nativeElement.focus();
  }

  editCode() {
    this.mainService.updateCode(this.solution);
  }

}
