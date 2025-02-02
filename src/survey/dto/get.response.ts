import { Field, ObjectType } from '@nestjs/graphql';
import { QuestionType } from './create.dto';

@ObjectType()
export class Question {
  @Field()
  id: number;

  @Field()
  title: string;

  @Field()
  type: QuestionType;

  @Field(() => [String])
  content: string[];
}

@ObjectType()
export class Survey {
  @Field()
  title: string;

  @Field()
  description: string;

  @Field(() => [Question])
  questions: Question[];
}
