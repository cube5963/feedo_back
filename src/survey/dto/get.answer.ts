import { ObjectType, Field, Int } from '@nestjs/graphql';

@ObjectType()
export class SurveyAns {
  @Field(() => Int)
  id: number;

  @Field()
  title: string;

  @Field({ nullable: true })
  description: string;

  @Field(() => [QuestionAns])
  questions: QuestionAns[];
}

@ObjectType()
export class QuestionAns {
  @Field(() => Int)
  id: number;

  @Field()
  title: string;

  @Field()
  type: string;

  @Field(() => [Response])
  responses: Response[];
}

@ObjectType()
export class Response {
  @Field(() => Int)
  id: number;

  @Field(() => [String])
  data: any[];
}
