export class Question {
  constructor(
    public surveyId: number,
    public title: string,
    public description: string,
    public order: number,
    public type: string,
    public questionType: QuestionType,
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
