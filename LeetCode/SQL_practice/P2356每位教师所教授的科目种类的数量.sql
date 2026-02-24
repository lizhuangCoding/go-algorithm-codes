
-- 力扣：https://leetcode.cn/problems/number-of-unique-subjects-taught-by-each-teacher/
-- 按 teacher_id 分组，然后计算不同的 subject_id

select teacher_id , count(DISTINCT subject_id)  as cnt
from Teacher
group by teacher_id
