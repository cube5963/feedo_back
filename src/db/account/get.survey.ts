import prisma from '../db.service';

export async function getSurvey(lgoinid: string): Promise<any> {
  const userid = await prisma.user.findUnique({
    where: {
      loginid: lgoinid,
    },
    select: {
      id: true,
    },
  });

  return await prisma.survey.findMany({
    where: {
      userId: userid.id,
    },
    select: {
      id: true,
      title: true,
    },
  });
}
