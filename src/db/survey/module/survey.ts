import { QuestionModule } from './question';

export class SurveyModule {
  constructor(
    public id: string,
    public title: string,
    public description: string,
    public userId: string,
    public questions: QuestionModule[],
  ) {}
}
