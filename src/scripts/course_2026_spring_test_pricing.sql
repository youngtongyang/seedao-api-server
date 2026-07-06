-- course_2026_spring 支付测试价：非会员 0.02 元，会员 0.01 元
-- 金额单位：分（fen）

BEGIN;

INSERT INTO courses (id, title, status, created_at, updated_at)
VALUES ('course_2026_spring', '2026 春季课', 'on_sale', NOW(), NOW())
ON CONFLICT (id) DO UPDATE SET
  title = EXCLUDED.title,
  status = EXCLUDED.status,
  updated_at = NOW();

INSERT INTO course_skus (course_id, sku_id, name, base_price_fen, status, created_at, updated_at)
VALUES ('course_2026_spring', 'standard', '标准票', 2, 'on_sale', NOW(), NOW())
ON CONFLICT (course_id, sku_id) DO UPDATE SET
  name = EXCLUDED.name,
  base_price_fen = EXCLUDED.base_price_fen,
  status = EXCLUDED.status,
  updated_at = NOW();

-- 软删除旧折扣规则（保留历史记录）
UPDATE course_discount_rules
SET status = 'off_sale', updated_at = NOW()
WHERE course_id = 'course_2026_spring' AND sku_id = 'standard' AND status = 'on_sale';

INSERT INTO course_discount_rules (
  course_id, sku_id, rule_type, discount_type, discount_value,
  status, created_at, updated_at
) VALUES (
  'course_2026_spring', 'standard', 'member', 'fixed', 1,
  'on_sale', NOW(), NOW()
);

COMMIT;

SELECT c.id, s.sku_id, s.base_price_fen AS non_member_fen,
       r.discount_value AS member_discount_fen,
       s.base_price_fen - r.discount_value AS member_final_fen
FROM courses c
JOIN course_skus s ON s.course_id = c.id
LEFT JOIN course_discount_rules r
  ON r.course_id = c.id AND r.sku_id = s.sku_id AND r.rule_type = 'member'
WHERE c.id = 'course_2026_spring';
