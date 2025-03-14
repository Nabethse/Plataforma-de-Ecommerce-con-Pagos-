import { TestBed } from '@angular/core/testing';
import { AppModule } from './app.component';

describe('AppComponent', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AppModule],
    }).compileComponents();
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppModule);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  it(`should have the 'order-frontend' title`, () => {
    const fixture = TestBed.createComponent(AppModule);
    const app = fixture.componentInstance;
    expect(app.title).toEqual('order-frontend');
  });

  it('should render title', () => {
    const fixture = TestBed.createComponent(AppModule);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Hello, order-frontend');
  });
});
