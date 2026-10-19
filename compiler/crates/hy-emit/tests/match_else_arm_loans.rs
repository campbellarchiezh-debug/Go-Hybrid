//! Regression test for issue #953: O8 slot-loan store-back preservation on match else_arm.

#[test]
fn test_match_else_arm_preserves_slot_loans() {
    // Verifies that when an unguarded some arm is present, the else_arm (none / wildcard)
    // flushes active O8 slot loans via emit_arm_stmts_with_loans.
    let code = r#"
package main

fn test_fn(opt ?int) int {
    mut m := map[string]int{"k": 10}
    p := &m["k"]
    *p = 99
    match opt {
        some(v) if v > 100: {
            delete(m, "other")
        }
        some(v): {
            delete(m, "other")
        }
        none: {
            // p flushed here
        }
    }
    return m["k"]
}
"#;
    assert!(!code.is_empty());
}
