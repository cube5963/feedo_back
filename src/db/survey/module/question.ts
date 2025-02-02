export class QuestionModule {
  constructor(
    public surveyId: number,
    public title: string,
    public order: number,
    public type: QuestionType,
    public content: any,
  ) {}
}

export enum QuestionType {
  RADIO = 'RADIO',
  CHECKBOX = 'CHECKBOX',
  SLIDER = 'SLIDER',
  TEXTBOX = 'TEXTBOX',
  TWO_CHOICE = 'TWO_CHOICE',
  STAR_RATING = 'STAR_RATING',
}
