import { Question } from './question';

export class Survey {
  constructor(
    public id: string,
    public title: string,
    public description: string,
    public userId: string,
    public questions: Question[],
  ) {}
}
