import { Mutation, Resolver, Args, Query, Int } from '@nestjs/graphql';
import { SurveyService } from './survey.service';
import { CreateSurveyInput } from './dto/create.dto';
import { Survey } from './dto/get.response';
import { AnswersInput } from './dto/answer.input';
import { SurveyAns } from './dto/get.answer';

@Resolver()
export class SurveyResolver {
  constructor(private readonly surveyService: SurveyService) {}

  @Query(() => Survey)
  async getSurvey(@Args('surveyid') surveyid: number) {
    return this.surveyService.Get(surveyid);
  }

  @Query(() => SurveyAns, { nullable: true })
  async getAnswer(
    @Args('surveyId', { type: () => Int }) surveyId: number,
  ): Promise<SurveyAns> {
    return await this.surveyService.GetAnswer(surveyId);
  }

  @Mutation(() => Int)
  async createSurvey(
    @Args('createsurveyinput') CreateSurveyInput: CreateSurveyInput,
  ): Promise<number> {
    return await this.surveyService.Create(CreateSurveyInput);
  }

  @Mutation(() => Boolean)
  async answerSurvey(
    @Args('answersurveyinput') answersInput: AnswersInput,
  ): Promise<boolean> {
    return await this.surveyService.Answer(answersInput);
  }
}
