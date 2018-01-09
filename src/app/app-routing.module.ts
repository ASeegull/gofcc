import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { Ch1Component } from './views/basic/ch1/ch1.component';
import { Ch2Component } from './views/basic/ch2/ch2.component';

const routes: Routes = [
  { path: 'ch1', component: Ch1Component },
  { path: 'ch2', component: Ch2Component }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
