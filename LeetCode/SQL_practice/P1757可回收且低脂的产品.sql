
-- 力扣：https://leetcode.cn/problems/recyclable-and-low-fat-products/description/

select p.product_id from Products as p where p.low_fats = 'Y' and p.recyclable = 'Y'
