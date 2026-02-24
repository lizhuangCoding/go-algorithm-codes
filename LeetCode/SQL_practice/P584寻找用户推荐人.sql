
-- 力扣：https://leetcode.cn/problems/find-customer-referee/?envType=study-plan-v2&envId=primers-list

select name
from Customer
where referee_id is null
   or referee_id != 2
