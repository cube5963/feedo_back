import { InputType, Field } from '@nestjs/graphql';

@InputType()
export class CreateSurveyInput {
  @Field()
  title: string;

  @Field()
  description: string;

  @Field()
  userId: string;

  @Field(() => [QuestionInput])
  questions: QuestionInput[];
}

@InputType()
export class QuestionInput {
  @Field()
  title: string;

  @Field()
  type: string;
  @Field(() => QuestionType)
  questionType: QuestionType;

  @Field()
  content: string;
}

export enum QuestionType {
  RADIO = 'RADIO',
  CHECKBOX = 'CHECKBOX',
  SLIDER = 'SLIDER',
  TEXTBOX = 'TEXTBOX',
  TWO_CHOICE = 'TWO_CHOICE',
  STAR_RATING = 'STAR_RATING',
}
