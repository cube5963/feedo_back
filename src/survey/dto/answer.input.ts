import { Field, InputType } from '@nestjs/graphql';

@InputType()
export class AnswerInput {
  @Field()
  questionId: number;

  @Field()
  answer: string;
}

@InputType()
export class AnswersInput {
  @Field()
  surveyId: number;

  @Field(() => [AnswerInput])
  answers: AnswerInput[];
}
